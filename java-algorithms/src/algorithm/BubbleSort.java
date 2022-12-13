package algorithm;

/**
 * time complexity O(n^2)
 */
public class BubbleSort {
    public void sort(int[] array) {
        for (var i = 0; i < array.length; i++) {
            var isSorted = true;
            for (var j = 0; j < array.length - 1; j++) {
                if (array[j + 1] < array[j]) {
                    swap(array, j);
                    isSorted = false;
                }
            }
            if (isSorted) {
                return;
            }
        }

    }

    public void swap(int[] array, int index) {
        var temp = array[index];
        array[index] = array[index + 1];
        array[index + 1] = temp;
    }

}
