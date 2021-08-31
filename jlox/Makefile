JFLAGS = -g
JC = javac
JVM = java

SRC = lox
MAIN = Lox

JAVA_FILES = $(wildcard $(SRC)/*.java)
CLASSES = $(JAVA_FILES:.java=.class)

default: $(CLASSES)

%.class: %.java
	$(JC) $(JFLAGS) $<

run: default 
	$(JVM) $(SRC)/$(MAIN)

clean:
	$(RM) $(SRC)/*.class
