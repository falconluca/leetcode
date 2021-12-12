package LRUCache;

import org.junit.Test;
import static org.junit.Assert.*;

public class LRUCacheTest {
    LRUCache lruCache;

    public LRUCacheTest() {
        lruCache = new LRUCache(2);
    }
    
    @Test
    public void test() {
        lruCache.put(1, 1);  // cache is {1=1}
        lruCache.put(2, 2);  // cache is {1=1, 2=2}
        assertEquals(1, lruCache.get(1));
        
        lruCache.put(3, 3);  // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
        assertEquals(-1, lruCache.get(2));
        
        lruCache.put(4, 4);  // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
        assertEquals(-1, lruCache.get(1));
        assertEquals(3, lruCache.get(3));
        assertEquals(4, lruCache.get(4));
    }
}
