package ImplementStrStr;

import org.junit.Test;

import static org.junit.Assert.*;

public class ImplementStrStrTest {
    private final ImplementStrStr implementStrStr;

    public ImplementStrStrTest() {
        this.implementStrStr = new ImplementStrStr();
    }

    @Test
    public void test() {
        assertEquals(2, implementStrStr.strStr("hello", "ll"));
        assertEquals(-1, implementStrStr.strStr("aaaaa", "bba"));
        assertEquals(0, implementStrStr.strStr("", ""));
    }
}
