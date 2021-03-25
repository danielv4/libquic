#include "quic.h"














void Multiaddr_func(char *sock_endpoint) 
{

	GoString endpoint = {sock_endpoint, strlen(sock_endpoint)};
	
	quic_multi_address(endpoint);
}


void GenerateECDSAKeyPair_func() 
{

	quic_generate_ecdsa_key_pair();
}


char *PeerID_func() 
{

	char *id = quic_peer_id();
	return id;
}


void Transport_func() 
{

	quic_transport();
}


void AcceptStream_func() 
{

	quic_accept_stream();
}


int StreamRead_func(char *bytes, int byte_size) 
{

	GoSlice go_bytes = {bytes, byte_size, byte_size};
	int val = quic_stream_read(go_bytes);
	return val;
}


int StreamWrite_func(char *bytes, int byte_size) 
{

	GoSlice go_bytes = {bytes, byte_size, byte_size};
	int val = quic_stream_write(go_bytes);
	return val;
}


void TransportConnect_func(char *peer_id) 
{

	GoString id = {peer_id, strlen(peer_id)};
	quic_transport_connect(id);
}


void OpenStream_func() 
{

	quic_open_stream();
}


void StreamReadChannel_func(void *func)
{

	pthread_t new_thread_id; 
    pthread_create(&new_thread_id, NULL, func, NULL);
}


struct QuicTransport NewQuicTransport()
{
    struct QuicTransport context;

	context.Multiaddr = Multiaddr_func;
	context.GenerateECDSAKeyPair = GenerateECDSAKeyPair_func;
	context.PeerID = PeerID_func;
	context.Transport = Transport_func;
	context.AcceptStream = AcceptStream_func;
	context.StreamRead = StreamRead_func;
	context.StreamWrite = StreamWrite_func;
	context.TransportConnect = TransportConnect_func;
	context.OpenStream = OpenStream_func;
	context.StreamReadChannel = StreamReadChannel_func;
	
	return context;
}











