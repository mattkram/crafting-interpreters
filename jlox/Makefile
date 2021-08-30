JFLAGS = -g
JC = javac
JVM = java
.SUFFIXES: .java .class
.java.class:
	$(JC) $(JFLAGS) $*.java

SRC = lox
MAIN = Lox

CLASSES = $(wildcard $(SRC)/*.java)

default: classes

classes: $(CLASSES:.java=.class)

run: lox/Lox.class
	$(JVM) $(SRC)/$(MAIN)

clean:
	$(RM) $(SRC)/*.class
