# nflogger

*nflogger* is a simple tool to dump packets logged via the Linux firewall NFLOG target.

Example usage:

```
sudo nft add ip filter INPUT log prefix \"filter/INPUT: \" group 100
sudo nflogger --group 100
```

When you are done press CTRL-C to stop nflogger and remove the rule again. To do this
print all rules of the chain INPUT like so.


```
sudo nft -a list chain filter INPUT
```

The output from above contains the handle ids of all rules. Lookup the id of your rule
and then delete it like so:

```
sudo nft delete rule filter INPUT handle <id>
```

## License

    3-clause BSD

    © 2023 Christian Pointner <equinox@spreadspace.org>
