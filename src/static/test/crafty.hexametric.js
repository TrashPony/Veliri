Crafty.extend({
    /**@
     * #Crafty.hexametric
     * @category 2D
     * Place entities in a hexagonal isometric fashion. It is similar to isometric but has another grid locations
     */
    hexametric: {
        _tile: {
            width: 0,
            height: 0,
            r: 0
        },
        _map: {
            width: 0,
            height: 0,
            x: 0,
            y: 0
        },

        _origin: {
            x: 0,
            y: 0
        },


        /**@
         * #Crafty.hexametric.init
         * @comp Crafty.hexametric
         * @sign public this Crafty.hexametric.init(Number tileWidth,Number tileHeight,Number mapWidth,Number mapHeight)
         * @param tileWidth - The size of base tile width in Pixel
         * @param tileHeight - The size of base tile height in Pixel
         * @param mapWidth - The width of whole map in Tiles
         * @param mapHeight - The height of whole map in Tiles
         *
         * Method used to initialize the size of the isometric placement.
         * Recommended to use a size alues in the power of `2` (128, 64 or 32).
         * This makes it easy to calculate positions and implement zooming.
         *
         * @example
         * ~~~
         * var iso = Crafty.hexametric.init(64,128,20,20);
         * ~~~
         *
         * @see Crafty.hexametric.place
         */
        init: function (tw, th, mw, mh) {
            this._tile.width = parseInt(tw, 10);
            this._tile.height = parseInt(th, 10) || parseInt(tw, 10) / 2;
            this._tile.r = this._tile.width / this._tile.height;

            this._map.width = parseInt(mw, 10);
            this._map.height = parseInt(mh, 10) || parseInt(mw, 10);

            this._origin.x = this._map.height * this._tile.width / 2;
            return this;
        },


        /**@
         * #Crafty.hexametric.place
         * @comp Crafty.hexametric
         * @sign public this Crafty.hexametric.place(Entity tile,Number x, Number y, Number layer)
         * @param tile - The entity that should be position in the isometric fashion
         * @param x - The `x` position to place the tile
         * @param y - The `y` position to place the tile
         * @param layer - The `z` position to place the tile (calculated by y position * layer)
         *
         * Use this method to place an entity in an isometric grid.
         *
         * @example
         * ~~~
         * var iso = Crafty.hexametric.init(64,128,20,20);
         * isos.place(Crafty.e('2D, DOM, Color').color('red').attr({w:128, h:128}),1,1,2);
         * ~~~
         *
         * @see Crafty.hexametric.size
         */
        place: function (obj, x, y, layer) {
            var pos = this.pos2px(x, y);
            if (!layer) layer = 1;
            var marginX = 0,
                marginY = 0;

//      if (obj.__margin !== undefined) {
//        marginX = obj.__margin[0];
//        marginY = obj.__margin[1];
//      }

            obj.x = pos.left + (marginX);
            obj.y = pos.top + (marginY);
            obj.z = (pos.top) * layer;
        },


        centerAt: function (x, y) {
            var pos = this.pos2px(x, y);
            Crafty.viewport.x = -pos.left + Crafty.viewport.width / 2 - this._tile.width;
            Crafty.viewport.y = -pos.top + Crafty.viewport.height / 2;
        },


        area: function (offset) {
            if (!offset) offset = 0;
            //calculate the corners
            var vp = Crafty.viewport.rect();
            var ow = offset * this._tile.width;
            var oh = offset * this._tile.height;
            vp._x -= (this._tile.width / 2 + ow);
            vp._y -= (this._tile.height / 2 + oh);
            vp._w += (this._tile.width / 2 + ow);
            vp._h += (this._tile.height / 2 + oh);
            /*  Crafty.viewport.x = -vp._x;
                Crafty.viewport.y = -vp._y;
                Crafty.viewport.width = vp._w;
                Crafty.viewport.height = vp._h;   */

            var grid = [];
            for (var y = vp._y, yl = (vp._y + vp._h); y < yl; y += this._tile.height / 2) {
                for (var x = vp._x, xl = (vp._x + vp._w); x < xl; x += this._tile.width / 2) {
                    var row = this.px2pos(x, y);
                    grid.push([~~row.x, ~~row.y]);
                }
            }
            return grid;
        },


        pos2px: function (x, y) {

            var left =  ( this._tile.width * x );
            if((x - 0.5) * (y&1))
                left = left - ( this._tile.width / 2 )

            var top = ( ( this._tile.height / 2 ) * y ) ;

            return {
                left: left,
                top: top
            };
        },


        px2pos: function (left, top) {

            var x, y;

            x = ( ( left / this._tile.width ) );
            y = ( top / ( this._tile.height / 2 ) );

            if(y%2)
                x += 0.5;

            return {
                x: x,
                y: y
            };
        },


        polygon: function (obj) {

            obj.requires("Collision");
            var marginX = 0,
                marginY = 0;
            if (obj.__margin !== undefined) {
                marginX = obj.__margin[0];
                marginY = obj.__margin[1];
            }
            var points = [
                [marginX - 0, obj.h - marginY - this._tile.height / 2],
                [marginX - this._tile.width / 2, obj.h - marginY - 0],
                [marginX - this._tile.width, obj.h - marginY - this._tile.height / 2],
                [marginX - this._tile.width / 2, obj.h - marginY - this._tile.height]
            ];
            var poly = new Crafty.polygon(points);
            return poly;

        }

    }
});
