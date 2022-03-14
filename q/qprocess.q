n:10000;
trade:([]time:"p"$.z.D-til n;sym:n?`a`b;price:n?10.0;qty:n?10);
getTrade:{select from trade where sym=x};
.z.pg:{[x]0N!(`zpg;x);value x};