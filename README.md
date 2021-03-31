基于go轻量级java虚拟机

## 已实现

- class 类加载器
- class 字节码解析
- runtime 运行时函数调用栈
- 指令解析：
    - 支持 const,push,ldc,load,store,pop 等堆栈指令
    - 支持 add,mul,div,sub 运算指令
    - 支持 icmp, ifeq,fcmpl,goto 等判断指令
    - 支持 get/setStatic,getField,SetFIeld 的类/实例字段指令
    - 支持 invokeVirtual,invokeStatic,invokeStatic...等函数调用指令
    - 支持 new,newArray 等对象创建指令
- 多态，支持类重载以及方法重写
- 方法区缓存，可以获取已加载的类实例

## 开始使用

要运行gvm可以使用命令行或者直接启动main()方法两种方式.

**命令行启动：**

1. 进入/src/share/ 路径下，直接执行

    ```go
    ./share -Xjre <jre path> -cp <class path> <class name>
    ```

    - -Xjre :  本地JDK中jre文件夹的位置（绝对路径），例如在Mac中，类似如下 `/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre`
    - -cp : class文件的绝对路径,提供给用户类加载器去查找class文件，以Mac为例，类似如下`/Users/zexho/project/gvm/java/src/`
    - class name : class文件名，不需要携带格式后缀（但是需要确保为.calss格式）

**从main()方法启动：**

1. 如果使用main()方法启动，同样需要保证传入启动参数，同样进入到/src/share路径下，然后执行,参数的定义与命令后启动相同。

    ```go
    go run main.go -Xjre <jre path> -cp <class path> <class name>
    ```

如果不想每次输入 -Xjre 和 -cp 参数，可以在`/src/share/launch_configuration` 中的JrePath和UserClassPath修改成你本地的值

## 展示

### 静态字段 & 实例字段 打印

```go
public class AlgorithmTest {
  static int i = 994;
  static float f = 994f;
  static double d = 994d;
  static boolean b = true;
  static long l = 994l;
  static String s = "hello world";

  public static void main(String[] args){
    GvmOut.to(i); 
    GvmOut.to(f); 
    GvmOut.to(d); 
    GvmOut.to(b); 
    GvmOut.to(l);
    GvmOut.to(s);
  }

}
```
![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e6ed8d5d-4ab2-4437-afbd-c67ab7e17c8e/Untitled.png](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/e6ed8d5d-4ab2-4437-afbd-c67ab7e17c8e/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210331%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210331T005901Z&X-Amz-Expires=86400&X-Amz-Signature=d33f1a70bb739e620862d07f8231a426182083cf4d3587eb8631bd651863ecaa&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

```go
public class PrintFieldsTest{
    public int i = 994;
    public float f = 994f;
    public double d = 994d;
    public boolean b = true;
    public long l = 994l;
    public String s = "hello world";
  
    public static void main(String[] args){
        PrintFieldsTest c = new PrintFieldsTest();
        GvmOut.to(c.i);
        GvmOut.to(c.f);
        GvmOut.to(c.d);
        GvmOut.to(c.b);
        GvmOut.to(c.l);
        GvmOut.to(c.s);
    }

}
```
![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c6557038-575b-4779-80a8-34f17d4198be/Untitled.png](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/c6557038-575b-4779-80a8-34f17d4198be/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210331%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210331T005940Z&X-Amz-Expires=86400&X-Amz-Signature=be8261c037d6178ed2ab21251516ec8cf7f53678e602f6eb2f9cf8a44361dd4c&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

### gvm 四则运算

```go
public class AlgorithmTest {
  
  static int i = 994;
  static float f = 994f;
  static double d = 994d;
  static boolean b = true;
  static long l = 994l;
  static String s = "hello world";

  public static void main(String[] args){
    GvmOut.to(i); 
    GvmOut.to(f); 
    GvmOut.to(d); 
    GvmOut.to(b); 
    GvmOut.to(l);
    GvmOut.to(s);

    AlgorithmTest.add();
    AlgorithmTest.div();
    AlgorithmTest.mul();
    AlgorithmTest.sub();
  }

  public static void add(){
    int i1 = 10;
    int i2 = 20;
    GvmOut.to(i1 + i2);

  }

  public static void div(){
    int i1 = 99;
    int i2 = 33;
    GvmOut.to(i1 / i2);
  }

  public static void mul(){
    int i1 = 100;
    int i2 = 4;
    GvmOut.to(i1 * i2);
  }

  public static void sub(){
    int i1 = 4;
    int i2 = 104;
    GvmOut.to(i1 - i2);
  }

}
```
![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/19415ef1-aa26-4038-ac7d-889895987251/Untitled.png](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/19415ef1-aa26-4038-ac7d-889895987251/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210331%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210331T010003Z&X-Amz-Expires=86400&X-Amz-Signature=223b3737cf29d5ecb26abab12a2ac40b4b4941bba424fb67e51cb15fc6e3d9ea&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

### 逻辑判断

`if`  & `while`  & `for`

```go
public static void if_test(){
      int n = 10;
      if (n > 9){
        GvmOut.to(10);
      } else{
        GvmOut.to(9);

      }
  }

  public static void for_test(){
    for(int i = 0 ; i < 5 ; i++){
      GvmOut.to(i);
    }

  }
  
  public static void while_test(){
    int x = 100;
    while(x < 105){
      GvmOut.to(x++);
    }

    x = 100;
    while(x < 105){
      GvmOut.to(++x);
    }
  }
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/14d46fc7-5779-45d1-88b2-21efa181820e/Untitled.png](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/14d46fc7-5779-45d1-88b2-21efa181820e/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210331%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210331T010019Z&X-Amz-Expires=86400&X-Amz-Signature=871200d8b94d9bfabcb8c4df8783d225adfd81e7f5edcf219ed85d21ae1063f1&X-Amz-SignedHeaders=host&response-content-disposition=filename%20%3D%22Untitled.png%22)

## 待实现

- [ ]  本地方法调用
- [ ]  异常堆栈
- [ ]  lambda语句支持
- [ ]  ⭐多线程
- [ ]  ⭐协程
- [ ]  ⭐锁：synchronized
- [ ]  ⭐反射
- [ ]  ⭐垃圾回收
- [ ]  ⭐gvm调优工具
