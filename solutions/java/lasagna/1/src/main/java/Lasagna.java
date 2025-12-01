public class Lasagna {
    // TODO: define the 'expectedMinutesInOven()' method

    public int expectedMinutesInOven () {
        return 40;
    }

    // TODO: define the 'remainingMinutesInOven()' method

    public int remainingMinutesInOven (int minutsInOven) {
        return expectedMinutesInOven() - minutsInOven;
    }


    // TODO: define the 'preparationTimeInMinutes()' method

    public int preparationTimeInMinutes (int layerAdded) {
        return layerAdded * 2; 
    }


    // TODO: define the 'totalTimeInMinutes()' method

    public int totalTimeInMinutes(int layerAdded, int min) {
        return preparationTimeInMinutes(layerAdded) + min;
    }

}
