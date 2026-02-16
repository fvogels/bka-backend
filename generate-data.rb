require 'date'


RANDOM = Random.new(4189)

def random_int(a, b)
  RANDOM.rand(b-a) + a
end

def random_pick(*xs)
  xs[random_int(0, xs.size)]
end

def pad(s, n)
  s.to_s.rjust(n, '0')
end

Header = Struct.new :bedrijfsnummer, :documentnummer, :boekjaar, :soort, :documentdatum, :boekingsdatum, :boekmaand, :invoerdatum, :invoertijd
Segment = Struct.new :bedrijfsnummer, :documentnummer, :boekjaar, :regelnummer, :identificatie, :vereffeningsdatum, :vereffeningsinvoerdatum, :vereffeningsdocument, :boekingssleutel

$headers = []
$segments = []


['1000', '1900', 'ABCD', 'AAAA'].each do |bedrijfsnummer|
  (2000..2025).each do |boekjaar|
    ndocs = random_int(2, 50)

    (1..ndocs).each do |k|
      documentnummer = k.to_s.rjust(10, '0')

      documentdatum = Date.new(boekjaar, random_int(1, 12), random_int(1, 29))
      boekingsdatum = documentdatum.next_day(random_int(0, 10))
      boekmaand = pad(boekingsdatum.month, 2)
      invoerdatum = boekingsdatum.next_day(random_int(0,5))

      hour = random_int(7, 19)
      minute = random_int(0, 60)
      second = random_int(0, 60)
      invoertijd = "#{pad(hour, 2)}#{pad(minute, 2)}#{pad(second, 2)}"

      soort = random_pick('KR', 'KZ', 'DR', 'DZ')
      $headers << Header.new(bedrijfsnummer, documentnummer, boekjaar, soort, documentdatum.strftime("%Y%m%d"), boekingsdatum.strftime("%Y%m%d"), boekmaand, invoerdatum.strftime("%Y%m%d"), invoertijd)

      nsegments = random_int(2, 10)
      (1..nsegments).each do |regelnummer|
        if random_int(0, 10) == 0
          vereffeningsdatum = '0' * 8
          vereffeningsinvoerdatum = '0' * 8
          vereffeningsdocument = '0' * 10
        else
          vereffeningsdatum = documentdatum.next_day(random_int(0, 30))
          vereffeningsinvoerdatum = vereffeningsdatum.next_day(random_int(0, 5))
          vereffeningsdocument = random_int(1, ndocs).to_s.rjust(10, '0')

          vereffeningsinvoerdatum = vereffeningsinvoerdatum.strftime("%Y%m%d")
          vereffeningsdatum = vereffeningsdatum.strftime("%Y%m%d")
        end

        boekingssleutel = random_pick('40', '50', '01', '11', '15', '21', '25', '31')

        $segments << Segment.new(bedrijfsnummer, documentnummer, boekjaar, pad(regelnummer, 3), 'X', vereffeningsdatum, vereffeningsinvoerdatum, vereffeningsdocument, boekingssleutel)
      end
    end
  end
end


abort "need two command line arguments" if ARGV.size != 2

File.open(ARGV[0], 'w') do |file|
  $headers.each do |header|
    file.puts "#{header.bedrijfsnummer},#{header.documentnummer},#{header.boekjaar},#{header.soort},#{header.documentdatum},#{header.boekingsdatum},#{header.boekmaand},#{header.invoerdatum},#{header.invoertijd}"
  end
end

File.open(ARGV[1], 'w') do |file|
  $segments.each do |segment|
    file.puts "#{segment.bedrijfsnummer},#{segment.documentnummer},#{segment.boekjaar},#{segment.regelnummer},#{segment.identificatie},#{segment.vereffeningsdatum},#{segment.vereffeningsinvoerdatum},#{segment.vereffeningsdocument},#{segment.boekingssleutel}"
  end
end