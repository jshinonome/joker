trade:([]time:"p"$.z.D-til 10;sym:10?`a`b;price:10?10.0;qty:10?10);
getTrade:{select from trade where sym=x};
.z.pg:{[x]0N!(`zpg;x);value x};