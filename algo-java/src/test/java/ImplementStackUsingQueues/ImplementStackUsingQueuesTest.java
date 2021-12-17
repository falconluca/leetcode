package ImplementStackUsingQueues;

import org.junit.Test;

import static org.junit.Assert.*;

public class ImplementStackUsingQueuesTest {
    private final ImplementStackUsingQueues stack;

    public ImplementStackUsingQueuesTest() {
        this.stack = new ImplementStackUsingQueues();
    }

    @Test
    public void test() {
        stack.push(1);
        stack.push(2);
        assertEquals(2, stack.top());
        assertEquals(2, stack.pop());
        assertFalse(stack.empty());
    }
}
