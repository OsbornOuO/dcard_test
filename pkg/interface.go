package pkg

// IService 抽象，請在這定義要實作的方法
type IService interface {
	IPRateLimitingService
}

// IRepository 抽象，請在這定義要實作的方法
type IRepository interface {
	IPRateLimitingRepository
}
