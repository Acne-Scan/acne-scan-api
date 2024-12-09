package conversion

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
)

func HistoryRequestToModel(request *web.HistoryRequest) *domain.History {
	return &domain.History{
		User_id:        request.User_id,
		Image:          domain.Images(request.Image),
		UserPicture:    request.UserPicture,
		ProductLink:    domain.ProductLink(request.ProductLink),
		Prediction:     request.Prediction,
		Recommendation: request.Recommendation,
	}
}
