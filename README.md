# Build

docker build -t open-exoplanet-catalogue .

# Run

docker run --rm -p 8080:8080 open-exoplanet-catalogue

# Deploy

helm install oec ./open-exoplanet-catalogue

# Output

```
Number of Orphan Planets i.e. if HostStarTempK is 0: 0
The Planet with the hottest star "V391 Peg b"
Planets discovered by year classified by Jupiter radii - small < 1, medium < 2, the rest large
key: 1781, value: oec.SizeCounts{small:1, medium:0, large:0}
key: 1846, value: oec.SizeCounts{small:1, medium:0, large:0}
key: 1930, value: oec.SizeCounts{small:1, medium:0, large:0}
key: 1992, value: oec.SizeCounts{small:4, medium:0, large:0}
key: 1994, value: oec.SizeCounts{small:1, medium:0, large:0}
key: 1995, value: oec.SizeCounts{small:1, medium:0, large:0}
key: 1996, value: oec.SizeCounts{small:6, medium:0, large:0}
key: 1997, value: oec.SizeCounts{small:1, medium:0, large:0}
key: 1998, value: oec.SizeCounts{small:4, medium:1, large:0}
key: 1999, value: oec.SizeCounts{small:7, medium:4, large:0}
key: 2000, value: oec.SizeCounts{small:17, medium:3, large:0}
key: 2001, value: oec.SizeCounts{small:11, medium:2, large:0}
key: 2002, value: oec.SizeCounts{small:23, medium:7, large:0}
key: 2003, value: oec.SizeCounts{small:23, medium:2, large:0}
key: 2004, value: oec.SizeCounts{small:23, medium:6, large:0}
key: 2005, value: oec.SizeCounts{small:28, medium:6, large:0}
key: 2006, value: oec.SizeCounts{small:22, medium:8, large:0}
key: 2007, value: oec.SizeCounts{small:41, medium:23, large:0}
key: 2008, value: oec.SizeCounts{small:34, medium:31, large:1}
key: 2009, value: oec.SizeCounts{small:64, medium:17, large:0}
key: 2010, value: oec.SizeCounts{small:73, medium:47, large:0}
key: 2011, value: oec.SizeCounts{small:121, medium:67, large:1}
key: 2012, value: oec.SizeCounts{small:100, medium:29, large:0}
key: 2013, value: oec.SizeCounts{small:98, medium:38, large:4}
key: 2014, value: oec.SizeCounts{small:888, medium:39, large:0}
key: 2015, value: oec.SizeCounts{small:153, medium:40, large:0}
key: 2016, value: oec.SizeCounts{small:1276, medium:28, large:0}
```