SumDivBy <- function(n,i)
{
  floor(floor(n/i)*(floor(n/i)+1)*i/2)
}

SumDivBy(999,3)+SumDivBy(999,5)-SumDivBy(999,15)
