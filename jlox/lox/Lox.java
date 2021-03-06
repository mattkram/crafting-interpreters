package lox;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class Lox {
    private static final Interpreter interpreter = new Interpreter();
    static boolean hadError = false;
    static boolean hadRuntimeError = false;

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
        if (hadRuntimeError) System.exit(70);
    }

    private static void runPrompt() throws IOException {
        // Run as an interactive prompt (REPL).
        InputStreamReader input = new InputStreamReader(System.in);
        BufferedReader reader = new BufferedReader(input);

        for (;;) {
            System.out.print(">>> ");
            String line = reader.readLine();
            if (line == null) break;
            try {
                run(line);
            } catch (RuntimeException error) {
            }
            hadError = false;
        }
    }

    private static void run(String source) {
        // Run the interpreter for a given source code string.
        Scanner scanner = new Scanner(source);
        List<Token> tokens = scanner.scanTokens();
        Parser parser = new Parser(tokens);
        List<Stmt> statements = parser.parse();

        // Stop if there was a syntax error.
        if (hadError) return;

        Resolver resolver = new Resolver(interpreter);
        resolver.resolve(statements);

        // Stop if there was a resolution error.
        if (hadError) return;

        //System.out.println(new AstPrinter().print(expression));
        interpreter.interpret(statements);
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

    static void error(Token token, String message) {
        if (token.type == TokenType.EOF) {
            report(token.line, " at end", message);
        } else {
            report(token.line, " at '" + token.lexeme + "'", message);
        }
    }

    static void runtimeError(RuntimeError error) {
        System.err.println(error.getMessage() +
            "\n[line " + error.token.line + "]");
        hadRuntimeError = true;
    }
}
