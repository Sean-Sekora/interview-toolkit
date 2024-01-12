package org.example.models;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Random;

public class Pond {

    public HashMap<Integer, LillyPad> lillyPadMap;

    public Pond(List<String> paths) {
        this.lillyPadMap = new HashMap<>();

        for (String path : paths) {
            String[] pathArray = path.split(",");
            Integer id = Integer.parseInt(pathArray[0]);
            this.lillyPadMap.put(id, new LillyPad(id, false, new ArrayList<>()));
        }

        for (String path : paths) {
            String[] pathArray = path.split(",");
            Integer id = Integer.parseInt(pathArray[0]);
            LillyPad lillyPad = this.lillyPadMap.get(id);
            for (int i = 1; i < pathArray.length; i++) {
                Integer pathId = Integer.parseInt(pathArray[i]);
                LillyPad pathLillyPad = this.lillyPadMap.get(pathId);
                lillyPad.paths.add(pathLillyPad);
            }
        }

    }

    public LillyPad placeFood(Integer foodLocationOverride) {
        LillyPad foodLocation = this.randomLillyPad(foodLocationOverride);
        foodLocation.isThereFood = true;
        return foodLocation;
    }

    public LillyPad birth(Integer birthPlaceOverride) {
        while (true) {
            LillyPad birthPlace = this.randomLillyPad(birthPlaceOverride);
            if (!birthPlace.isThereFood) {
                return birthPlace;
            }
        }
    }

    private LillyPad randomLillyPad(Integer id) {
        if (id != null) {
            return this.lillyPadMap.get(id);
        }
        Random rand = new Random();
        int randomIndex = rand.nextInt(this.lillyPadMap.size()-1) + 1;
        return this.lillyPadMap.get(randomIndex);
    }
}