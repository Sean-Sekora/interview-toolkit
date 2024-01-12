package org.example;

import org.example.models.LillyPad;
import org.example.models.Pond;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    private final List<String> paths = Arrays.asList(
        "1,2,4",
        "2,1,4",
        "3,6,9,10",
        "4,1,2,5,6",
        "5,4,7",
        "6,3,4,10",
        "7,5,10",
        "8,3,9",
        "9,3,8",
        "10,6,7"
    );

    @Test
    void findFood() {
        Pond pond = new Pond(paths);
        Solution solution = new Solution(pond);
        solution.pond.placeFood(null);
        LillyPad birthplace = solution.pond.birth(null);
        LillyPad foodLocation = solution.findFood(birthplace, null);
        assertNotNull(foodLocation);
    }
}