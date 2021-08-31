package lox;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class Lox {
    static boolean hadError = false;

    public static void main(String[] args) throws IOException {
        // Main program entry point, processing CLI args and delegating to 
        // different run modes.
        if (args.length == 2 && args[0].equals("-c")) {
            runString(args[1]);
        } else if (args.length > 1) {
            System.out.println("Usage: jlox [script]");
            System.exit(64);
        } else if (args.length == 1) {
            runFile(args[0]);
        } else {
            runPrompt();
        }
    }

    private static void runString(String source) throws IOException {
        // Run a string passed in with the -c argument.
        run(source);
        
        if (hadError) System.exit(65);
    }

    private static void runFile(String path) throws IOException {
        // Read in a file and run the interpreter.
        byte[] bytes = Files.readAllBytes(Paths.get(path));
        run(new String(bytes, Charset.defaultCharset()));

        // Indicate an error in the exit code.
        if (hadError) System.exit(65);
    }

    private static void runPrompt() throws IOException {
        // Run as an interactive prompt (REPL).
        InputStreamReader input = new InputStreamReader(System.in);
        BufferedReader reader = new BufferedReader(input);

        for (;;) {
            System.out.print("> ");
            String line = reader.readLine();
            if (line == null) break;
            run(line);
            hadError = false;
        }
    }

    private static void run(String source) {
        // Run the interpreter for a given source code string.
        Scanner scanner = new Scanner(source);
        List<Token> tokens = scanner.scanTokens();

        // For now, just print the tokens.
        for (Token token : tokens) {
            System.out.println(token);
        }
    }

    static void error(int line, String message) {
        // Emit an error message to stderr.
        report(line, "", message);
    }

    private static void report(int line, String where, String message) {
        // Report an error by printing the line and message.
        System.err.println(
            "[line " + line + "] Error" + where + ": " + message);
        hadError = true;
    }
}
