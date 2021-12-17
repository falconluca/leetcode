package ImplementStackUsingQueues;

import java.util.LinkedList;
import java.util.Queue;

/**
 * 225. Implement Stack using Queues
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 */
public class ImplementStackUsingQueues {
    private final Queue<Integer> q1;
    private final Queue<Integer> q2;
    private int size;
    private boolean isSymmetric;

    public ImplementStackUsingQueues() {
        this.q1 = new LinkedList<>();
        this.q2 = new LinkedList<>();
        this.size = 0;
        this.isSymmetric = false;
    }

    public void push(int el) {
        if (isSymmetric) {
            q1.offer(el);
        }
        else {
            q2.offer(el);
        }
        size++;
    }
    
    public int pop() {
        int result;
        if (isSymmetric) {
            while (q1.size() > 1) {
                q2.offer(q1.poll());
            }
            result = q1.poll();
        }
        else {
            while (q2.size() > 1) {
                q1.offer(q2.poll());
            }
            result = q2.poll();
        }
        isSymmetric = !isSymmetric;
        size--;
        return result;
    }
    
    public int top() {
        int result;
        if (isSymmetric) {
            while (q1.size() > 1) {
                q2.offer(q1.poll());
            }
            result = q1.poll();
            q2.offer(result);
        }
        else {
            while (q2.size() > 1) {
                q1.offer(q2.poll());
            }
            result = q2.poll();
            q1.offer(result);
        }
        isSymmetric = !isSymmetric;
        return result;
    }
    
    public boolean empty() {
        return size == 0;
    }
}
