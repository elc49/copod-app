package com.lomolo.copodv2.ui.common

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedCard
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.lomolo.copodv2.R
import com.lomolo.copodv2.ui.screens.Land
import com.lomolo.copodv2.ui.theme.CopodV2Theme

@Composable
fun LandCard(
    modifier: Modifier = Modifier,
    land: Land,
) {
    OutlinedCard(
        Modifier.wrapContentHeight()
    ) {
        Image(
            painterResource(R.drawable.sealed),
            contentScale = ContentScale.Fit,
            modifier = Modifier
                .fillMaxWidth()
                .size(200.dp),
            contentDescription = null,
        )
        Box(
            Modifier
                .background(MaterialTheme.colorScheme.secondaryContainer)
                .fillMaxWidth()
                .padding(12.dp),
        ) {
            Row {
                Column {
                    Text(
                        land.titleNo,
                        fontWeight = FontWeight.Bold,
                    )
                    Text(land.town)
                }
                Spacer(Modifier.weight(1f))
                Text(
                    "${land.size} ${land.symbol}",
                    style = MaterialTheme.typography.titleLarge,
                )
            }
        }
    }
}

@Composable
fun NoLands(
    modifier: Modifier = Modifier,
) {
    Column(
        Modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Icon(
            painterResource(R.drawable.sealed),
            modifier = Modifier.size(60.dp),
            contentDescription = stringResource(R.string.land),
        )
        Text(stringResource(R.string.no_lands))
    }
}