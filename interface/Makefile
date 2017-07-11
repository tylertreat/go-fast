BARGRAPH := bargraph
SIZE := 700

%.png: %.ppm
		mogrify -reverse -flatten $<
			mogrify -resize ${SIZE}x${SIZE} -format png $<
%.ppm: %.perf
		$(BARGRAPH) -fig $< | fig2dev -L ppm -m 4 > $@

clean:
	-rm -f *.png
