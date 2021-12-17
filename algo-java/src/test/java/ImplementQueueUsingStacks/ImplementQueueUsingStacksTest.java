package ImplementQueueUsingStacks;

import org.junit.Test;

import static org.junit.Assert.*;

public class ImplementQueueUsingStacksTest {
    private final ImplementQueueUsingStacks queue;

    public ImplementQueueUsingStacksTest() {
        this.queue = new ImplementQueueUsingStacks();
    }

    @Test
    public void test() {
        queue.push(1);
        queue.push(2);
        assertEquals(1, queue.peek());
        assertEquals(1, queue.pop());
        assertFalse(queue.empty());
    }
}
