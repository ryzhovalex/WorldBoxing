# Orwynn protocol v2: Making reactivity.
Keypoints:
* we have a thread per connection for reading input
* we have a single thread for output queue, but a separate thread per context sending
* each subscriber is called in a separate thread
* we don't manage inner messages: all out messages are going outside (target connection must be set). It's no more possible to exchange messages within inner host. This is the sacrifice we made for optimization.
