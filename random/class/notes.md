## explicit
Specifies that a constructor or conversion function (since C++11) or deduction guide (since C++17) is explicit, that is, it cannot be used for implicit conversions and copy-initialization.

#### implicit conversions

An implicit conversion sequence is the sequence of conversions required to convert an argument in a function call to the type of the corresponding parameter in a function declaration. The compiler tries to determine an implicit conversion sequence for each argument.

The compiler tries to determine an implicit conversion sequence for each argument. It then categorizes each implicit conversion sequence in one of three categories and ranks them depending on the category. The compiler does not allow any program in which it cannot find an implicit conversion sequence for an argument.

The following are the three categories of conversion sequences in order from best to worst:
- Standard conversion sequences
- User-defined conversion sequences
- Ellipsis conversion sequences

#### copy initialization
The Copy initialization can be done using the concept of copy constructor. As we know that the constructors are used to initialize the objects. We can create our copy constructor to make a copy of some other object, or in other words, initialize current object with the value of another object. On the other hand, the direct initialization can be done using assignment operation.

The main difference between these two types of initialization is that the copy initialization creates a separate memory block for the new object. But the direct initialization does not make new memory space. It uses reference variable to point to the previous memory block.
```
classname (const classname &obj) {  
  // body of constructor 
}

classname Ob1, Ob2; 
Ob2 = Ob1;
```