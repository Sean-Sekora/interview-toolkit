package org.example;

import org.example.models.LillyPad;
import org.example.models.Pond;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Solution {

    private static final List<String> paths = Arrays.asList(
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

    public final Pond pond;

    public Solution(Pond pond) {
        this.pond = pond;
    }

    public static void main(String[] args) {
        // Create the pond
        Solution solution = new Solution(new Pond(paths));

        // Place the food
        LillyPad foodLocation = solution.pond.placeFood(null);
        System.out.println("Food Location: " + foodLocation.id);

        // Pick the Frog's birthplace
        LillyPad birthplace = solution.pond.birth(null);

        // Find food
        LillyPad foundFood = solution.findFood(birthplace, null);
        System.out.println("Found food at Lilly Pad: " + foundFood.id);
    }

    public LillyPad findFood(LillyPad frogsLocation, Set<LillyPad> visited) {
        if (visited == null) {
            visited = new HashSet<>();
        }
        visited.add(frogsLocation);

        if (frogsLocation.isThereFood) {
            return frogsLocation;
        }
        System.out.println("The frog is on lilly pad " + frogsLocation.id);

        for (LillyPad lillyPad: frogsLocation.paths) {
            if (!visited.contains(lillyPad)) {
                return findFood(lillyPad, visited);
            }
        }

        // TODO: Implement this method
        return null;
    }
}