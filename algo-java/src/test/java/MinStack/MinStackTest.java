package MinStack;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class MinStackTest {
    MinStack minStack;

    public MinStackTest() {
        this.minStack = new MinStack();
    }

    @Test
    public void test() {
        minStack.push(-2);
        minStack.push(0);
        minStack.push(-3);
        assertEquals(-3, minStack.getMin());
        minStack.pop();
        assertEquals(0, minStack.top());
        assertEquals(-2, minStack.getMin());
    }
}
