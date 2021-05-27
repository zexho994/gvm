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

1. 在项目根路径下，直接执行

    ```go
    ./share -Xjre <jre path> -cp <class path> <class name>
    ```

    - -Xjre :  本地JDK中jre文件夹的位置（绝对路径），例如在Mac中，类似如下 `/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre`
    - -cp : class文件的绝对路径,提供给用户类加载器去查找class文件，以Mac为例，类似如下`/Users/zexho/project/gvm/java/src/`
    - class name : class文件名，不需要携带格式后缀（但是需要确保为.calss格式）

**从main()方法启动：**

1. 如果使用main()方法启动，同样需要保证传入启动参数，同样进入到项目根路径下，然后执行,参数的定义与命令后启动相同。

    ```go
    go run main.go -Xjre <jre path> -cp <class path> <class name>
    ```

如果不想每次输入 -Xjre 和 -cp 参数，可以在`/src/share/launch_configuration` 中的JrePath和UserClassPath修改成你本地的值

## 展示

### 静态字段 & 实例字段 打印

```java
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
输出

![pic1](https://tva1.sinaimg.cn/large/008eGmZEly1gphbp71970j307l03zt8y.jpg)

```java
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
输出

![pic1](https://tva1.sinaimg.cn/large/008eGmZEly1gphbr3nuvhj307o046glu.jpg)

### gvm 四则运算

```java
public class AlgorithmTest {

  public static void main(String[] args){
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
输出

![pic3](https://tva1.sinaimg.cn/large/008eGmZEly1gphbrw78bij305402p3yk.jpg)

### 逻辑判断

`if`  & `while`  & `for`

```java
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
输出

![pic4](https://tva1.sinaimg.cn/large/008eGmZEly1gphbtozlsxj304w0avwf5.jpg)

## 待实现

- [ ]  实现 jni
- [ ]  异常堆栈
- [ ]  lambda 语句支持
- [ ]  多线程
- [ ]  协程
- [ ]  锁：synchronized
- [ ]  反射
- [ ]  gc
