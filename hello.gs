let fib = fn(a) {
  if (a == 1) {
    return 1;
  };
  return a + fib(a-1)
};

fib(100000);


