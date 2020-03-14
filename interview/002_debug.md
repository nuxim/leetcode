## 线上服务 CPU 很高该怎么做？有哪些措施可以找到问题
   
   定位出现问题的堆栈信息排查具体问题
   
   1、top命令：Linux命令。可以查看实时的CPU使用情况。也可以查看最近一段时间的CPU使用情况。
   
   2、ps命令： Linux命令。强大的进程状态监控命令。可以查看进程以及进程中线程的当前CPU使用情况。属于当前状态的采样数据。
   
   3、jstack：  Java提供的命令。可以查看某个进程的当前线程栈运行情况。根据这个命令的输出可以定位某个进程的所有线程的当前运行状态、运行代码，以及是否死锁等等。
   
   4、pstack：Linux命令。可以查看某个进程的当前线程栈运行情况
##
##