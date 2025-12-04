import java.io.File;

import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;

public class Sample {

    public static void main(String args[]) throws Exception {
        Sample s = new Sample();
        System.out.println("Working directory: " + System.getProperty("user.dir"));
        List<String> stringList = Files.readAllLines(Paths.get("input.txt"));
        int[][] matrix = s.computeMatrixFromInput(stringList);
        var computed = s.computeNeighbour(matrix);
        s.printMatrix(computed);
        int count = 0;
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                if (matrix[i][j] > 0 && computed[i][j] < 5) {
                    count++;
                }
            }
        }
        System.out.println("\n" + count);
    }

    public void printMatrix(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            System.out.println("\n-------------------------------------------------------------------------------");
            for (int j = 0; j < matrix[i].length; j++) {
                System.out.print(matrix[i][j] + "");
            }
        }
    }

    public int[][] computeMatrixFromInput(List<String> stringList) {
        int[][] matrix = new int[stringList.size()][stringList.get(0).length()];
        for (int i = 0; i < stringList.size(); i++) {
            var charArr = stringList.get(i).toCharArray();
            for (int j = 0; j < charArr.length; j++) {
                matrix[i][j] = charArr[j] == '@' ? 1 : 0;
            }
        }
        return matrix;
    }

    public int[][] computeNeighbour(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                matrix = hasNeighbour(matrix, i, j);
            }
        }
        return matrix;
    }

    public boolean isLeftColumnAvailable(int j) {
        return j > 0;
    }

    public boolean isRightColumnAvailable(int j, int colLength) {
        return j < colLength - 1;
    }

    public boolean isBottomRowAvailable(int i, int rowLength) {
        return i < rowLength - 1;
    }

    public boolean isTopRowAvailable(int i) {
        return i > 0;
    }

    public int[][] hasNeighbour(int[][] matrix, int i, int j) {
        if (matrix[i][j] <= 0) {
            return matrix;
        }
        if (isTopRowAvailable(i)) {
            if (isLeftColumnAvailable(j)) {
                matrix[i][j] += matrix[i - 1][j - 1] > 0 ? 1 : 0;
            }
            if (isRightColumnAvailable(j, matrix[0].length)) {
                matrix[i][j] += matrix[i - 1][j + 1] > 0 ? 1 : 0;
            }
            matrix[i][j] += matrix[i - 1][j] > 0 ? 1 : 0;

        }
        if (isBottomRowAvailable(i, matrix.length)) {
            if (isLeftColumnAvailable(j)) {
                matrix[i][j] += matrix[i + 1][j - 1] > 0 ? 1 : 0;
            }
            if (isRightColumnAvailable(j, matrix[0].length)) {
                matrix[i][j] += matrix[i + 1][j + 1] > 0 ? 1 : 0;
            }
            matrix[i][j] += matrix[i + 1][j] > 0 ? 1 : 0;
        }
        if (isLeftColumnAvailable(j)) {
            matrix[i][j] += matrix[i][j - 1] > 0 ? 1 : 0;
        }
        if (isRightColumnAvailable(j, matrix[0].length)) {
            matrix[i][j] += matrix[i][j + 1] > 0 ? 1 : 0;
        }

        return matrix;
    }

}