% Code adapted from github.com/thomastc, because someone insisted that that didn't work in matlab proper ;)
handle = fopen('puzzle13.in', 'r');
input = reshape(fscanf(handle, '%d: %d'), 2, []);
fclose(handle);
depths = input(1, :);
ranges = input(2, :);
total_depth = depths(length(depths));
severities = depths .* ranges;
delay = 0;
while sum(mod(delay + depths, 2 * ranges - 2) == 0) > 0
  delay = delay + 1;
end;
disp(delay);