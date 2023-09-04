type Pool struct{
	Register 	chan *Client
	Unregister 	chan *Client
	Clients		mac[*Client]bool
	Broadcast chan Message
}


NewPool()