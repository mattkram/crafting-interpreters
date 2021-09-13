JFLAGS = -g
JC = javac
JVM = java

SRC = lox
MAIN = Lox

JAVA_FILES = $(wildcard $(SRC)/*.java)
CLASSES = $(JAVA_FILES:.java=.class)

default: ast $(CLASSES)

%.class: %.java
	$(JC) $(JFLAGS) $<

run: default 
	$(JVM) $(SRC)/$(MAIN)

ast: tool/GenerateAst.class
	java tool/GenerateAst $(shell pwd)/lox

clean:
	$(RM) $(SRC)/*.class tool/*.class
	$(RM) $(SRC)/Expr.java $(SRC)/Stmt.java
