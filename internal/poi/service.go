package poi

type repo interface {
	Create(poi *POI) error
	List() ([]POI, error)
}

type service struct {
	repo repo
}

func NewService(rp repo) *service {
	return &service{
		repo: rp,
	}
}

func (svc *service) RegisterNewPOI(poi *POI) error {
	return svc.repo.Create(poi)
}

func (svc *service) GetAllPOI() ([]POI, error) {
	return svc.repo.List()
}

func (svc *service) GetAllPOIByDistance(refPoint *point, distanceMax uint64) ([]POI, error) {
	pois, err := svc.repo.List()
	nearerPOIs := make([]POI, 0, len(pois))
	if err != nil {
		return nil, err
	}
	for _, poi := range pois {
		if poi.Distance(refPoint) <= float64(distanceMax) {
			nearerPOIs = append(nearerPOIs, poi)
		}
	}
	return nearerPOIs, nil
}
