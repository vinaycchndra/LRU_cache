## Written a Learning Project For The Caching in Golang with lru and lfu eviction strategies.
    # LRU Cache Key Notes.
        -Least Recently Used Cache in Golang. A least recently used cache stores the order of items in a most recently used order or the items which are not accessed very recently will be pushed to the last in the queue.

        - We will use a doubly linked list data structure to mantain the order of the cached values.
                                                                    
        - We will use hash map to access the cached value in O(1) time.

        - We will use another object(struct) to access the head and tail of the doubly linked list and matains the maximum value of the value of the cache value.


        Process: 
        -We will initialize the Cache having the Queue and hashMap. The Que will be initialized with the head and tail with the nil values. These only two nodes also initialize our doubly linked 	list.

        - When we fetch a value from the cache the value becomes the recently use it should we removed and should be added to the begining of the linked list and if we exceeded the length of the cache the node value at the last should be removed since it is least fetched.

        - If we are accessing the values which are not precent in the doubly linked list will not alter the structure of the cache.

        - If we are adding anything to the cache that value should also be added at the begining of the queue.

    # LFU Cache Key Notes.
        - In Lfu cache we mantain a count of frequency of every key to check how frequently it has been accessed.
        - Once the cache is full the key with the least frequently count will be evicted.
        - Here we have written a min heap data structure based implementation of the cache.
        - So, Popping the min heap would pop the key with the minimum frequency.
        - Here, we are also using the timestamp feature which is just an integer that signifies the current time so when an item is accessed, 
          inserted or updated the current time stamp is being stamped on the key item. This timestamp we can use to pop the item with the older time stamp in case of the tie in the frequency.
        - We mantain a cache struct with the key and value as pointer to the item of the key. The min heap also maintains the pointer of the  
          cached item.
        - The item object in itself holds the key, its frequency, the index at which it resides in the heap so at the time of popping a lfu item 
          using the key we can delete it from the cache. Also, index is used to point from which node we should start heapifying in the min heap when there is any update in the frequency or value. So while heapifying also we need to keep in the sync the index key in the item.