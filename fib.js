let fib = function(a) {
  if (a == 1) return 1;
  return a * fib(a-1)
}

console.log(fib(20))