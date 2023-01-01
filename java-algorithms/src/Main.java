import algorithm.BubbleSort;
import algorithm.InsertionSort;
import algorithm.MergeSort;
import algorithm.SelectionSort;

import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        int[] numbers = {3, 3, 1 ,54 ,6, 3, 2};
//        var sorter = new BubbleSort();
//        var sorter = new SelectionSort();
//        var sorter = new InsertionSort();
        var sorter = new MergeSort();
        sorter.sort(numbers);
        System.out.println(Arrays.toString(numbers));
    }
}