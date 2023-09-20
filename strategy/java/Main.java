// Main.java

public class Main
{
    public static void main(String[] args) {
        Dice dice = new Dice();
        
        // some sort of logic to decide the strategy
        // could also be passed into a constructor instead of a setter
        RollStrategy rollStrat = new FairRollStrategy();
        // RollStrategy rollStrat = new RiggedRollStrategy(); // alternate strat
        dice.setStrategy(rollStrat);
        
        dice.showRoll();
    }
}
