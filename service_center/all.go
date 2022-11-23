package service_center

func InitAllService() error {
	// init restful service
	for _, srv := range restfulServices {
		err := srv.Config()
		if err != nil {
			return err
		}
	}
	// init grpc service
	return nil
}
