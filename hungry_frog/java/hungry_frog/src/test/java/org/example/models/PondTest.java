package org.example.models;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class PondTest {
    private final List<String> paths = Arrays.asList(
        "1,2,3",
        "2,1,3",
        "3,1,2"
    );

    @Test
    void NewPond() {
        Pond pond = new Pond(paths);
        pond.placeFood(1);
        assertEquals(3, pond.lillyPadMap.size());
        assertEquals(pond.lillyPadMap.get(2), pond.lillyPadMap.get(1).paths.get(0));
        assertEquals(pond.lillyPadMap.get(3), pond.lillyPadMap.get(1).paths.get(1));
        assertEquals(pond.lillyPadMap.get(1), pond.lillyPadMap.get(2).paths.get(0));
        assertEquals(pond.lillyPadMap.get(3), pond.lillyPadMap.get(2).paths.get(1));
        assertEquals(pond.lillyPadMap.get(1), pond.lillyPadMap.get(3).paths.get(0));
        assertEquals(pond.lillyPadMap.get(2), pond.lillyPadMap.get(3).paths.get(1));
        assertTrue(pond.lillyPadMap.get(1).isThereFood);
        assertFalse(pond.lillyPadMap.get(2).isThereFood);
        assertFalse(pond.lillyPadMap.get(3).isThereFood);
    }
}