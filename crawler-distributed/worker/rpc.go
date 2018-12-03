package worker

import "crawler/engine"

type CrawlService struct {}

func (CrawlService) Process(req Request, result *ParseResult) error {
	enginerReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineRequest, err := engine.Worker(enginerReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineRequest)
	return nil
}