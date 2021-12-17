package ImplementStrStr;

/**
 * 28. Implement strStr()
 * https://leetcode.com/problems/implement-strstr/
 */
public class ImplementStrStr {
    public int strStr(String haystack, String needle) {
        if (haystack == null || needle == null) {
            return -1;
        }
        
        int shouldCheckHaystackIdxRange = haystack.length() 
                // 仅需检查匹配串的一个字符
                - needle.length() + 1;
        for (int haystackIdx = 0; haystackIdx < shouldCheckHaystackIdxRange; haystackIdx++) {
            int needleIdx;
            for (needleIdx = 0; needleIdx < needle.length(); needleIdx++) {
                int ptr = haystackIdx + needleIdx;
                if (haystack.charAt(ptr) != needle.charAt(needleIdx)) {
                    // 对不上就检查haystack的下一个char
                    break;
                }
            }
            if (needleIdx == needle.length()) { 
                // 整个needle都有检查完, 并且没有被break中断
                return haystackIdx;
            }
        }
        
        // Not found
        return -1;
    }
}
