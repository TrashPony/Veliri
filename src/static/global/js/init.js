function InitGame() {
    var game = new Phaser.Game(800, 600, Phaser.AUTO, 'map', null, true, true);

    var BasicGame = function (game) { };

    BasicGame.Boot = function (game) { };

    var isoGroup, cursorPos;

    BasicGame.Boot.prototype =
        {
            preload: function () {
                game.load.image('tile', './assets/tile.png');
                game.load.image('castle', './assets/castle.png');


                game.time.advancedTiming = true;

                // Добавляем и включаем плагин
                game.plugins.add(new Phaser.Plugin.Isometric(game));

                // Это используется для установки смещения на холсте, в игре для изометрической координаты. 0, 0, 0 - по умолчанию
                // эта точка была бы на координатах экрана 0, 0 (вверху слева), что обычно нежелательно.
                game.iso.anchor.setTo(0.5, 0.2);


            },
            create: function () {

                // создаем группу для наших спрайтов
                isoGroup = game.add.group();

                // Загрушает плитки на сетке
                this.spawnTiles();

                // Обеспечивает 3д положение курсора
                cursorPos = new Phaser.Plugin.Isometric.Point3();
            },
            update: function () {
                // Обновите положение крсора
                // Важно понимать, что экран-изометрическая проекция означает, что вам нужно указать позицию z вручную, поскольку это нелегко
                // определяемый из позиции двумерного указателя без дополнительных обманщиков. По умолчанию позиция z равна 0, если не установлена.
                game.iso.unproject(game.input.activePointer.position, cursorPos);

                // Прокрутите все плитки и проверьте, совпадает ли 3D-положение сверху с автоматически созданной границей плитки IsoSprite.
                isoGroup.forEach(function (tile) {
                    var inBounds = tile.isoBounds.containsXY(cursorPos.x, cursorPos.y);
                    // Если да, сделайте небольшую анимацию и оттенок.
                    if (!tile.selected && inBounds) {

                        console.log(tile.gridPosition.x + ":" + tile.gridPosition.y);

                        tile.selected = true;
                        tile.tint = 0x86bfda;
                        game.add.tween(tile).to({ isoZ: 4 }, 200, Phaser.Easing.Quadratic.InOut, true);
                    }
                    // Если нет, вернитесь к тому, как это было.
                    else if (tile.selected && !inBounds) {
                        tile.selected = false;
                        tile.tint = 0xffffff;
                        game.add.tween(tile).to({ isoZ: 0 }, 200, Phaser.Easing.Quadratic.InOut, true);
                    }
                });
            },
            render: function () {
                game.debug.text(game.time.fps || '--', 2, 14, "#a7aebe");
            },
            spawnTiles: function () {
                var tile;
                var size = 10;
                for (var x = 0; x < size; x += 1) {
                    for (var y = 0; y < size; y += 1) {
                        var xx = x * 38;
                        var yy = y * 38;

                        // Создайте спрайт, используя новый метод game.add.isoSprite в указанной позиции.
                        // Последний параметр - это группа, которую вы хотите добавить (как и game.add.sprite)
                        if (x === 5 && y === 5) {
                            tile = game.add.isoSprite(xx, yy, 0, 'castle', 0, isoGroup);
                        } else {
                            tile = game.add.isoSprite(xx, yy, 0, 'tile', 0, isoGroup);
                        }

                        // добавляем свойство позиции на координатной сети
                        tile.gridPosition = {};
                        tile.gridPosition.x = x;
                        tile.gridPosition.y = y;


                        tile.anchor.set(0.5, 0);
                    }
                }
            }
        };

    game.state.add('Boot', BasicGame.Boot);
    game.state.start('Boot');
}