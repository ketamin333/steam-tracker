package pricehistoryservice

import "context"

func (s *Service) Run(ctx context.Context) error {
	tg, err := s.trackedGameRepo.GetAll(ctx)

	if err != nil {
		return err
	}

}
