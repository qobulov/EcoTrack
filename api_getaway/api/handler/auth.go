package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Ro'yxatdan o'tish funktsiyasi
func register(c *gin.Context) {
	// Ro'yxatdan o'tish logikasini shu yerga yozing
	c.JSON(http.StatusOK, gin.H{"message": "Ro'yxatdan o'tish muvaffaqiyatli!"})
}

// Kirish funktsiyasi
func login(c *gin.Context) {
	// Kirish logikasini shu yerga yozing
	c.JSON(http.StatusOK, gin.H{"message": "Kirish muvaffaqiyatli!"})
}

// Chiqish funktsiyasi
func logout(c *gin.Context) {
	// Chiqish logikasini shu yerga yozing
	c.JSON(http.StatusOK, gin.H{"message": "Chiqish muvaffaqiyatli!"})
}

// Tokenni yangilash funktsiyasi
func refreshToken(c *gin.Context) {
	// Tokenni yangilash logikasini shu yerga yozing
	c.JSON(http.StatusOK, gin.H{"message": "Token yangilandi!"})
}

// Rate limiting Middleware
func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Rate limiting logikasini shu yerga yozing
		// Misol uchun, bu yerda har bir foydalanuvchi 1 daqiqada faqat 5 ta so'rov yuborishi mumkin
		rateLimit := 5
		interval := time.Minute

		ip := c.ClientIP()
		count, exists := requestCount[ip]

		if exists && count >= rateLimit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Juda ko'p so'rovlar. Iltimos, keyinroq qayta urinib ko'ring."})
			return
		}

		if !exists {
			requestCount[ip] = 1
		} else {
			requestCount[ip]++
		}

		go resetRequestCount(ip, interval)
		c.Next()
	}
}

// So'rovlar sonini qayta tiklash funksiyasi
func resetRequestCount(ip string, interval time.Duration) {
	time.Sleep(interval)
	delete(requestCount, ip)
}
