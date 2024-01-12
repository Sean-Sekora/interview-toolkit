package org.example.models;

import java.util.List;

public class LillyPad {
    public Integer id;
    public Boolean isThereFood;
    public List<LillyPad> paths;

    public LillyPad(Integer id, Boolean isThereFood, List<LillyPad> paths) {
        this.id = id;
        this.isThereFood = isThereFood;
        this.paths = paths;
    }
}