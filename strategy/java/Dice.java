//Dice.java
public class Dice{
    private RollStrategy rollStrategy;
    
    public void setStrategy(RollStrategy rollStrat){
        this.rollStrategy = rollStrat;
    }
    
    public void showRoll(){
        System.out.println("rolled a: " + rollStrategy.roll());
    }
}
