
/**
 * @author Zexho
 * @date 2021/1/8 7:31 下午
 */
@FunctionalInterface
public interface Soap<R, P> {
    R to(P... p); 
}
