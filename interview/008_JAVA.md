## HashMap 的查询时间复杂度
   　　理想情况下是 O(1)的，但是实际中会出现 hash 碰撞，导致无法达到效果
## LinkedList和ArrayList的区别
* LinkedList 底层是基于双向链表实现的，而 ArrayList 底层是基于动态数组实现的；
* 查询的时候 LinkedList 的效率要低于 ArrayList，因为 LinkedList 需要遍历链表，而 ArrayList 底层数组根据下标直接获取数据。
* 插入删除数据的时候，LinkedList 效率比ArrayList 效率高，因为 ArrayList 在数据多的情况下会进行数组扩容或移动数组。
## ThreadLocal的使用场景
   　　ThreadLocal 适用于每个线程需要自己独立的实例且该实例需要在多个方法中被使用，也即变量在线程间隔离而在方法或类间共享的场景
## 堆内存和栈内存有什么区别
* 堆内存是线程共享的，栈内存是线程私有的；
* 堆内存用来存放由new创建的对象和数组，栈内存中存放一些基本类型的变量和对象的引用变量
## 反射的机制
   
   大家都知道，要让Java程序能够运行，那么就得让Java类要被Java虚拟机加载。Java类如果不被Java虚拟机加载，是不能正常运行的。现在我们运行的所有的程序都是在编译期的时候就已经知道了你所需要的那个类的已经被加载了。
   
   Java的反射机制是在编译并不确定是哪个类被加载了，而是在程序运行的时候才加载、探知、自审。使用在编译期并不知道的类。这样的特点就是反射
   
   反射机制通过void setAccessible(boolean flag)方法可以得到一个类的private的方法和属性，使用这些private的方法和属性
## Set 和 List 区别？
   
   Set（集）：集合中的对象不按特定方式排序，并且没有重复对象。它的有些实现类能对集合中的对象按特定方式排序。
   
   List（列表）：集合中的对象按索引位置排序，可以有重复对象，允许按照对象在集合中的索引位置检索对象。
## ArrayList 和 LinkedList 区别
   
    ArrayList是实现了基于动态数组的数据结构，LinkedList基于链表的数据结构
   
   ArrayList 继承AbstractList
   LinkedList 继承AbstractSequentialList
   ArrayList 采用的是数组形式来保存对象的，这种方式将对象放在连续的位置中，所以最大的缺点就是插入删除时非常麻烦
   
   LinkedList 采用的将对象存放在独立的空间中，而且在每个空间中还保存下一个链接的索引 但是缺点就是查找非常麻烦 要丛第一个索引开始
## 如果存取相同的数据，ArrayList 和 LinkedList 谁占用空间更大？
   
   对于随机访问get和set，ArrayList觉得优于LinkedList，因为LinkedList要移动指针
   
   对于新增和删除操作add和remove，LinedList比较占优势，因为ArrayList要移动数据，若要从数组中删除或插入某一个对象，需要移动后段的数组元素，从而会重新调整索引顺序,调整索引顺序会消耗一定的时间，相反,LinkedList是使用链表实现的,若要从链表中删除或插入某一个对象,只需要改变前后对象的引用即可
## Concurrenthashmap 是怎么做到线程安全的？
   
   ConcurrentHashMap的大部分操作和HashMap是相同的，例如初始化，扩容和链表向红黑树的转变等。但是，在ConcurrentHashMap中，大量使用了U.compareAndSwapXXX
   
   的方法，这个方法是利用一个CAS算法实现无锁化的修改值的操作，他可以大大降低锁代理的性能消耗。这个算法的基本思想就是不断地去比较当前内存中的变量值与你指定的
   
   一个变量值是否相等，如果相等，则接受你指定的修改的值，否则拒绝你的操作。因为当前线程中的值已经不是最新的值，你的修改很可能会覆盖掉其他线程修改的结果。这一
   
   点与乐观锁，SVN的思想是比较类似的。
   
   同时，在ConcurrentHashMap中还定义了三个原子操作，用于对指定位置的节点进行操作。这三种原子操作被广泛的使用在ConcurrentHashMap的get和put等方法中，
   
   正是这些原子操作保证了ConcurrentHashMap的线程安全。
   
   在ConcurrentHashMap没有出现以前，jdk使用hashtable来实现线程安全，但是hashtable是将整个hash表锁住，所以效率很低下。
   
   ConcurrentHashMap将数据分别放到多个Segment中，默认16个，每一个Segment中又包含了多个HashEntry列表数组，
   
   对于一个key，需要经过三次hash操作，才能最终定位这个元素的位置，这三次hash分别为：
   
   对于一个key，先进行一次hash操作，得到hash值h1，也即h1 = hash1(key)；
   将得到的h1的高几位进行第二次hash，得到hash值h2，也即h2 = hash2(h1高几位)，通过h2能够确定该元素的放在哪个Segment；
   将得到的h1进行第三次hash，得到hash值h3，也即h3 = hash3(h1)，通过h3能够确定该元素放置在哪个HashEntry。
   每一个Segment都拥有一个锁，当进行写操作时，只需要锁定一个Segment，而其它Segment中的数据是可以访问的。
   
## HashTable 你了解过吗？
   
   Hashtable既不支持Null key也不支持Null value。Hashtable的put()方法的注释中有说明
   
   Hashtable是线程安全的，
   
   Hashtable是线程安全的，它的每个方法中都加入了Synchronize方法，效率比较低
   
   Hashtable默认的初始大小为11，之后每次扩充，容量变为原来的2n+1。
   
   Hashtable在计算元素的位置时需要进行一次除法运算，而除法运算是比较耗时的
## synchronized、lock
   
   synchronized是java中的一个关键字，也就是说是Java语言内置的特性
   
   如果一个代码块被synchronized修饰了，当一个线程获取了对应的锁，并执行该代码块时，其他线程便只能一直等待，等待获取锁的线程释放锁，而这里获取锁的线程释放锁只会有两种情况：
   
   　　1）获取锁的线程执行完了该代码块，然后线程释放对锁的占有；
   
   　　2）线程执行发生异常，此时JVM会让线程自动释放锁
   
   那么如果这个获取锁的线程由于要等待IO或者其他原因（比如调用sleep方法）被阻塞了，但是又没有释放锁，其他线程便只能干巴巴地等待，试想一下，这多么影响程序执行效率。
   
   　　因此就需要有一种机制可以不让等待的线程一直无期限地等待下去（比如只等待一定的时间或者能够响应中断），通过Lock就可以办到
   
   再举个例子：当有多个线程读写文件时，读操作和写操作会发生冲突现象，写操作和写操作会发生冲突现象，但是读操作和读操作不会发生冲突现象。
   
   　　但是采用synchronized关键字来实现同步的话，就会导致一个问题：
   
   　　如果多个线程都只是进行读操作，所以当一个线程在进行读操作时，其他线程只能等待无法进行读操作。
   
   　　因此就需要一种机制来使得多个线程都只是进行读操作时，线程之间不会发生冲突，通过Lock就可以办到。
   
   另外，通过Lock可以知道线程有没有成功获取到锁。这个是synchronized无法办到的
## lock 和 synchronized 的区别？
   
       1）Lock是一个接口，而synchronized是Java中的关键字，synchronized是内置的语言实现；
   
   　　2）synchronized在发生异常时，会自动释放线程占有的锁，因此不会导致死锁现象发生；而Lock在发生异常时，如果没有主动通过unLock()去释放锁，则很可能造成死锁现象，因此使用Lock时需要在finally块中释放锁；
   
   　　3）Lock可以让等待锁的线程响应中断，而synchronized却不行，使用synchronized时，等待的线程会一直等待下去，不能够响应中断；
   
   　　4）通过Lock可以知道有没有成功获取锁，而synchronized却无法办到。
   
   　　5）Lock可以提高多个线程进行读操作的效率。
   
   　　在性能上来说，如果竞争资源不激烈，两者的性能是差不多的，而当竞争资源非常激烈时（即有大量线程同时竞争），此时Lock的性能要远远优于synchronized。所以说，在具体使用时要根据适当情况选择
## 读写锁设计主要解决什么问题？
   
   多线程，
   
   读操作可以共享，写操作是排他的，读可以有多个在读，写只有唯一个在写，同时写的时候不允许读
   
   解决了读和读可以同时进行，读和写不能同时进行，写和写不能同时进行
##
##
##