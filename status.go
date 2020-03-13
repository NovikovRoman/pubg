package pubg

func (c Client) Status() bool {
	_, statusCode, err := c.requestGET(EmptyPlatform, "/status")
	return err == nil && statusCode == 200
}
