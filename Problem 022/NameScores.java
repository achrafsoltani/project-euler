import java.io.*;
import java.util.*;

public class NameScores {

    public static void main(String[] args) {

        try {
            File names = new File("0022_names.txt");
            Scanner sc = new Scanner(names);
            String data = sc.nextLine();
            //System.out.println(data);
            data = data.replace("\"", "");
            String[] arrOfStr = data.split(",", 0);
            Arrays.sort(arrOfStr);
            int pos = 1;
            long sum = 0;
            for (String s : arrOfStr) {
                //System.out.println(s);
                sum += pos * score(s);
                pos++;
            }
            System.out.println(sum);

        } catch (Exception e){
            System.out.println(e.getMessage());
        }

    }

    public static int score(String name){
        String[] chars = name.split("(?!^)");
        int score = 0;
        for (String c : chars) {
            score += ((int)c.charAt(0) - 64);
        }
        return score;
    }

}