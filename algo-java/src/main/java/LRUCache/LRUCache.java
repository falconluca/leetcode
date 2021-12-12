package LRUCache;

import java.util.Iterator;
import java.util.LinkedHashMap;

/**
 * 146. LRU Cache
 * https://leetcode.com/problems/lru-cache/
 */
public class LRUCache {
    private int capacity;
    
    private final LinkedHashMap<Integer, Integer> map;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.map = new LinkedHashMap<>();
    }
    
    public int get(int key) {
        Integer value = map.get(key);
        if (value == null) {
            return -1;
        }

        put(key, value);
        return value;
    }
    
    public void put(int key, int value) {
        if (map.size() == this.capacity) {
            Iterator<Integer> iter = map.keySet().iterator();
            iter.next();
            iter.remove();
        }
        
        map.put(key, value);
    }
}
