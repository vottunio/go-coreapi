package coreapi

func (c *CoreApi) SendDynamicRequest(s *SendCoreTransactionModel) error {

	return c.sendCoreTransaction(s)

}
