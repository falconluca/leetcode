package MinStack;

import java.util.Stack;

/**
 * 155. Min Stack
 * https://leetcode.com/problems/min-stack/description/
 */
public class MinStack {
    private final Stack<Integer> s1;
    private final Stack<Integer> s2;

    public MinStack() {
        s1 = new Stack<>();
        s2 = new Stack<>();
    }

    public void push(int val) {
        if (s2.isEmpty()) {
            s2.push(val);
        }
        else {
            int min = Math.min(val, getMin());
            s2.push(min);
        }
        s1.push(val);
        
    }

    public void pop() {
        s1.pop();
        s2.pop();
    }

    public int top() {
        return s1.peek();
    }

    public int getMin() {
        return s2.peek();
    }
}
