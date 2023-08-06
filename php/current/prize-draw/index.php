<?php




    echo(rank("Addison,Jayden,Sofia,Michael,Andrew,Lily,Benjamin", [4, 2, 1, 4, 3, 1, 2], 4));

    function rank($st, $we, $n) {
        $participants = [];
        $list = explode(',', $st);

        if ($st == '') {
            return 'No participants';
        }

        if ($n > count($list)) {
            return 'Not enough participants';
        }

        for ($i = 0; $i < count($list); $i++) {
            $val = strlen($list[$i]);
            $chars = str_split($list[$i]);
            foreach ($chars as $char) {
                $val = $val + alphaDig(strtoupper($char));
            }
            $new = new stdClass;
            $new->name = $list[$i];
            $new->score = $val * $we[$i];
            array_push($participants, $new);
        }

        usort($participants, function($a, $b){return $a->score > $b->score; });
        $res = $participants[$n-1]->name;

        foreach($participants as $p) {
            if ($participants[$n-1]->score == $p->score) {
                $res = function($a,$b){return $a->name > $b->name; };
            }
        }
        return $res;
    };

    function alphaDig($char) {
        $alphabet = range('A', 'Z');
        return array_search($char, $alphabet);
    };
?>