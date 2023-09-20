// FairRollStrategy.java
import java.util.Random;

public class FairRollStrategy implements RollStrategy {
    public int roll(){
        Random rand = new Random(); 
        return rand.nextInt(5)+1;
    }
}
